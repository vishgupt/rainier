use tonic::{Request, Response, Status};
use crate::core::collection::{InMemoryCollectionManager, CollectionManager};
use crate::proto::vectordb::{
    collection_service_server::{CollectionService, CollectionServiceServer},
    ListCollectionsRequest, ListCollectionsResponse,
    GetCollectionRequest, GetCollectionResponse,
};

#[derive(Debug)]
pub struct VectorCollectionService {
    collection_manager: InMemoryCollectionManager,
}

impl VectorCollectionService {
    pub fn new() -> Self {
        Self {
            collection_manager: CollectionManager::new(),
        }
    }
    
    pub fn server() -> CollectionServiceServer<VectorCollectionService> {
        CollectionServiceServer::new(Self::new())
    }
}

#[tonic::async_trait]
impl CollectionService for VectorCollectionService {
    async fn list_collections(
        &self,
        request: Request<ListCollectionsRequest>,
    ) -> Result<Response<ListCollectionsResponse>, Status> {
        let req = request.into_inner();
        let collections = self.collection_manager.list_collections(&req.database_name);
        
        let proto_collections = collections;
        
        let total = proto_collections.len() as i32;
        let response = ListCollectionsResponse {
            collections: proto_collections,
            total,
        };
        
        Ok(Response::new(response))
    }
    
    async fn get_collection(
        &self,
        request: Request<GetCollectionRequest>,
    ) -> Result<Response<GetCollectionResponse>, Status> {
        let req = request.into_inner();
        
        match self.collection_manager.get_collection(&req.database_name, &req.collection_name) {
            Some(col) => {
                let response = GetCollectionResponse {
                    collection: Some(col.clone()),
                };
                
                Ok(Response::new(response))
            }
            None => Err(Status::not_found(format!(
                "Collection '{}' not found in database '{}'",
                req.collection_name, req.database_name
            ))),
        }
    }
}
