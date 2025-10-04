use serde::{Deserialize, Serialize};
use std::collections::HashMap;

// Collection-related types
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Collection {
    pub name: String,
    pub dimension: usize,
    pub metric: DistanceMetric,
    pub vectors_count: usize,
    pub index_config: IndexConfig,
    pub status: CollectionStatus,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateCollectionRequest {
    pub name: String,
    pub dimension: usize,
    pub metric: DistanceMetric,
    pub index_config: Option<IndexConfig>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CollectionResponse {
    pub collections: Vec<Collection>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum CollectionStatus {
    Creating,
    Ready,
    Error,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum DistanceMetric {
    Cosine,
    Euclidean,
    DotProduct,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct IndexConfig {
    pub index_type: IndexType,
    pub m: Option<usize>,           // HNSW parameter
    pub ef_construct: Option<usize>, // HNSW parameter
}

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum IndexType {
    Hnsw,
    Flat,
}

// Vector-related types
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Vector {
    pub id: String,
    pub values: Vec<f32>,
    pub metadata: Option<HashMap<String, serde_json::Value>>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpsertRequest {
    pub vectors: Vec<Vector>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpsertResponse {
    pub upserted_count: usize,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct GetVectorsRequest {
    pub ids: Vec<String>,
    pub include_values: Option<bool>,
    pub include_metadata: Option<bool>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct GetVectorsResponse {
    pub vectors: Vec<Vector>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DeleteVectorsRequest {
    pub ids: Option<Vec<String>>,
    pub filter: Option<HashMap<String, serde_json::Value>>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DeleteVectorsResponse {
    pub deleted_count: usize,
}

// Search-related types
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SearchRequest {
    pub vector: Vec<f32>,
    pub top_k: usize,
    pub filter: Option<HashMap<String, serde_json::Value>>,
    pub include_values: Option<bool>,
    pub include_metadata: Option<bool>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SearchMatch {
    pub id: String,
    pub score: f32,
    pub values: Option<Vec<f32>>,
    pub metadata: Option<HashMap<String, serde_json::Value>>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SearchResponse {
    pub matches: Vec<SearchMatch>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct BatchSearchRequest {
    pub queries: Vec<SearchRequest>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct BatchSearchResponse {
    pub results: Vec<SearchResponse>,
}

// Error types
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ErrorResponse {
    pub error: String,
    pub message: String,
}

// Default implementations
impl Default for IndexConfig {
    fn default() -> Self {
        Self {
            index_type: IndexType::Hnsw,
            m: Some(16),
            ef_construct: Some(200),
        }
    }
}
