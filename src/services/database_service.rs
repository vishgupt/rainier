use tonic::{Request, Response, Status};
use crate::core::database::{InMemoryDatabaseManager, DatabaseManager};
use crate::proto::vectordb::{
    database_service_server::{DatabaseService, DatabaseServiceServer},
    ListDatabasesRequest, ListDatabasesResponse,
    GetDatabaseRequest, GetDatabaseResponse,
    Database,
};

#[derive(Debug)]
pub struct VectorDatabaseService {
    db_manager: InMemoryDatabaseManager,
}

impl VectorDatabaseService {
    pub fn new() -> Self {
        Self {
            db_manager: InMemoryDatabaseManager::new(),
        }
    }
    
    pub fn server() -> DatabaseServiceServer<VectorDatabaseService> {
        DatabaseServiceServer::new(Self::new())
    }
}

#[tonic::async_trait]
impl DatabaseService for VectorDatabaseService {
    async fn list_databases(
        &self,
        _request: Request<ListDatabasesRequest>,
    ) -> Result<Response<ListDatabasesResponse>, Status> {
        let databases = self.db_manager.list_databases();
        
        let proto_databases: Vec<Database> = databases
            .into_iter()
            .map(|db| Database {
                name: db.name,
                description: db.description,
                created_at: db.created_at,
                updated_at: db.updated_at,
                metadata: db.metadata,
            })
            .collect();
        
        let total = proto_databases.len() as i32;
        let response = ListDatabasesResponse {
            databases: proto_databases,
            total,
        };
        
        Ok(Response::new(response))
    }
    
    async fn get_database(
        &self,
        request: Request<GetDatabaseRequest>,
    ) -> Result<Response<GetDatabaseResponse>, Status> {
        let req = request.into_inner();
        
        match self.db_manager.get_database(&req.database_name) {
            Some(db) => {
                let proto_db = Database {
                    name: db.name.clone(),
                    description: db.description.clone(),
                    created_at: db.created_at,
                    updated_at: db.updated_at,
                    metadata: db.metadata.clone(),
                };
                
                let response = GetDatabaseResponse {
                    database: Some(proto_db),
                };
                
                Ok(Response::new(response))
            }
            None => Err(Status::not_found(format!(
                "Database '{}' not found",
                req.database_name
            ))),
        }
    }
}
