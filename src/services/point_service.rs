use tonic::{Request, Response, Status};
use crate::core::point::{InMemoryPointManager, PointManager};
use crate::proto::vectordb::{
    point_service_server::{PointService, PointServiceServer},
    GetPointRequest, GetPointResponse,
    SearchNearestRequest, SearchNearestResponse,
    PointWithDistance,
};

#[derive(Debug)]
pub struct VectorPointService {
    point_manager: InMemoryPointManager,
}

impl VectorPointService {
    pub fn new() -> Self {
        Self {
            point_manager: PointManager::new(),
        }
    }
    
    pub fn server() -> PointServiceServer<VectorPointService> {
        PointServiceServer::new(Self::new())
    }
}

#[tonic::async_trait]
impl PointService for VectorPointService {
    async fn get_point(
        &self,
        request: Request<GetPointRequest>,
    ) -> Result<Response<GetPointResponse>, Status> {
        let req = request.into_inner();
        
        match self.point_manager.get_point(&req.database_name, &req.collection_name, &req.point_id) {
            Some(point) => {
                let response = GetPointResponse {
                    point: Some(point.clone()),
                };
                
                Ok(Response::new(response))
            }
            None => Err(Status::not_found(format!(
                "Point '{}' not found in collection '{}' in database '{}'",
                req.point_id, req.collection_name, req.database_name
            ))),
        }
    }
    
    async fn search_nearest(
        &self,
        request: Request<SearchNearestRequest>,
    ) -> Result<Response<SearchNearestResponse>, Status> {
        let req = request.into_inner();
        
        if req.query_vector.is_empty() {
            return Err(Status::invalid_argument("Query vector cannot be empty"));
        }
        
        let limit = if req.limit > 0 { req.limit as usize } else { 10 };
        
        let results = self.point_manager.search_nearest(
            &req.database_name,
            &req.collection_name,
            &req.query_vector,
            limit,
        );
        
        let proto_results: Vec<PointWithDistance> = results
            .into_iter()
            .map(|(point, distance)| PointWithDistance {
                point: Some(point),
                distance,
            })
            .collect();
        
        let response = SearchNearestResponse {
            results: proto_results,
        };
        
        Ok(Response::new(response))
    }
}
