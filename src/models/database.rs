use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Database {
    pub name: String,
    pub description: String,
    pub created_at: i64,
    pub updated_at: i64,
    pub metadata: HashMap<String, String>,
}

impl Database {
    pub fn new(name: String, description: String) -> Self {
        let now = chrono::Utc::now().timestamp();
        Self {
            name,
            description,
            created_at: now,
            updated_at: now,
            metadata: HashMap::new(),
        }
    }
}
