use crate::proto::vectordb::Database;
use std::collections::HashMap;
use super::database_manager::DatabaseManager;

#[derive(Debug)]
pub struct InMemoryDatabaseManager {
    databases: HashMap<String, Database>,
}

impl InMemoryDatabaseManager {
    fn create_database(name: String, description: String) -> Database {
        let now = chrono::Utc::now().timestamp();
        Database {
            name,
            description,
            created_at: now,
            updated_at: now,
            metadata: std::collections::HashMap::new(),
        }
    }
}

impl DatabaseManager for InMemoryDatabaseManager {
    fn new() -> Self {
        let mut manager = Self {
            databases: HashMap::new(),
        };
        
        // Add a sample database for testing
        let sample_db = Self::create_database(
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
    
    fn add_database(&mut self, database: Database) -> Result<(), String> {
        // Check if database already exists
        if self.databases.contains_key(&database.name) {
            return Err(format!("Database '{}' already exists", database.name));
        }
        
        // Add the database
        self.databases.insert(database.name.clone(), database);
        Ok(())
    }
}
