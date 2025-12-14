use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Collection {
    pub name: String,
    pub database_name: String,
    pub description: String,
    pub vector_dimension: i32,
    pub created_at: i64,
    pub updated_at: i64,
    pub metadata: HashMap<String, String>,
}

impl Collection {
    pub fn new(
        name: String,
        database_name: String,
        description: String,
        vector_dimension: i32,
    ) -> Self {
        let now = chrono::Utc::now().timestamp();
        Self {
            name,
            database_name,
            description,
            vector_dimension,
            created_at: now,
            updated_at: now,
            metadata: HashMap::new(),
        }
    }
}
