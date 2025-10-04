use axum::{
    routing::{delete, get, post},
    Router,
};
use tower::ServiceBuilder;
use tower_http::cors::CorsLayer;

use crate::handlers::*;

pub fn create_router() -> Router {
    Router::new()
        // Health check
        .route("/health", get(health_check))
        
        // Collection management routes
        .route("/collections", post(create_collection))
        .route("/collections", get(list_collections))
        .route("/collections/:collection_name", get(get_collection))
        .route("/collections/:collection_name", delete(delete_collection))
        
        // Vector operations routes
        .route("/collections/:collection_name/vectors/upsert", post(upsert_vectors))
        .route("/collections/:collection_name/vectors", get(get_vectors))
        .route("/collections/:collection_name/vectors", delete(delete_vectors))
        
        // Search routes
        .route("/collections/:collection_name/search", post(search_vectors))
        .route("/collections/:collection_name/search/batch", post(batch_search_vectors))
        
        // Add middleware
        .layer(
            ServiceBuilder::new()
                .layer(CorsLayer::permissive()) // Enable CORS for web clients
                .into_inner(),
        )
}
