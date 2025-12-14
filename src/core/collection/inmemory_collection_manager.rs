use crate::proto::vectordb::Collection;
use std::collections::HashMap;
use super::collection_manager::CollectionManager;

#[derive(Debug)]
pub struct InMemoryCollectionManager {
    collections: HashMap<String, Vec<Collection>>, // database_name -> collections
}

impl InMemoryCollectionManager {
    fn create_collection(
        name: String,
        database_name: String,
        description: String,
        vector_dimension: i32,
    ) -> Collection {
        let now = chrono::Utc::now().timestamp();
        Collection {
            name,
            database_name,
            description,
            vector_dimension,
            created_at: now,
            updated_at: now,
            metadata: std::collections::HashMap::new(),
        }
    }
}

impl CollectionManager for InMemoryCollectionManager {
    fn new() -> Self {
        let mut manager = Self {
            collections: HashMap::new(),
        };
        
        // Add sample collections for testing
        let default_collections = vec![
            Self::create_collection(
                "products".to_string(),
                "default".to_string(),
                "Product embeddings".to_string(),
                768,
            ),
            Self::create_collection(
                "users".to_string(),
                "default".to_string(),
                "User embeddings".to_string(),
                512,
            ),
        ];
        
        manager.collections.insert("default".to_string(), default_collections);
        
        manager
    }
    
    fn list_collections(&self, database_name: &str) -> Vec<Collection> {
        self.collections
            .get(database_name)
            .map(|cols| cols.clone())
            .unwrap_or_default()
    }
    
    fn get_collection(&self, database_name: &str, collection_name: &str) -> Option<&Collection> {
        self.collections
            .get(database_name)
            .and_then(|cols| cols.iter().find(|c| c.name == collection_name))
    }
    
    fn add_collection(&mut self, collection: Collection) -> Result<(), String> {
        // Validate that the collection's database exists
        if !self.collections.contains_key(&collection.database_name) {
            return Err(format!("Database '{}' does not exist", collection.database_name));
        }
        
        // Check if collection already exists
        if let Some(collections) = self.collections.get(&collection.database_name) {
            if collections.iter().any(|c| c.name == collection.name) {
                return Err(format!(
                    "Collection '{}' already exists in database '{}'",
                    collection.name, collection.database_name
                ));
            }
        }
        
        // Add the collection
        self.collections
            .get_mut(&collection.database_name)
            .unwrap()
            .push(collection);
        
        Ok(())
    }
}
