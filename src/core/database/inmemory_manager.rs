use crate::models::database::Database;
use std::collections::HashMap;
use super::manager::DatabaseManager;

#[derive(Debug)]
pub struct InMemoryDatabaseManager {
    databases: HashMap<String, Database>,
}

impl DatabaseManager for InMemoryDatabaseManager {
    fn new() -> Self {
        let mut manager = Self {
            databases: HashMap::new(),
        };
        
        // Add a sample database for testing
        let sample_db = Database::new(
            "default".to_string(),
            "Default vector database".to_string(),
        );
        manager.databases.insert("default".to_string(), sample_db);
        
        manager
    }
    
    fn list_databases(&self) -> Vec<Database> {
        self.databases.values().cloned().collect()
    }
    
    fn get_database(&self, name: &str) -> Option<&Database> {
        self.databases.get(name)
    }
}
