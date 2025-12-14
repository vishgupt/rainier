mod models;
mod services;
mod core;
mod api;
mod storage;

use crate::proto::vectordb::database_service_server::DatabaseServiceServer;
use crate::proto::vectordb::collection_service_server::CollectionServiceServer;
use tonic::transport::Server;

pub mod proto {
    pub mod vectordb {
        tonic::include_proto!("vectordb");
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "127.0.0.1:50051".parse()?;
    
    println!("Vector Database gRPC Server listening on {}", addr);
    
    Server::builder()
        .add_service(DatabaseServiceServer::new(services::database_service::VectorDatabaseService::new()))
        .add_service(CollectionServiceServer::new(services::collection_service::VectorCollectionService::new()))
        .serve(addr)
        .await?;
    
    Ok(())
}
