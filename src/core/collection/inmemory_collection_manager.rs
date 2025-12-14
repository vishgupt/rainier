use crate::models::collection::Collection;
use std::collections::HashMap;
use super::collection_manager::CollectionManager;

#[derive(Debug)]
pub struct InMemoryCollectionManager {
    collections: HashMap<String, Vec<Collection>>, // database_name -> collections
}

impl CollectionManager for InMemoryCollectionManager {
    fn new() -> Self {
        let mut manager = Self {
            collections: HashMap::new(),
        };
        
        // Add sample collections for testing
        let default_collections = vec![
            Collection::new(
                "products".to_string(),
                "default".to_string(),
                "Product embeddings".to_string(),
                768,
            ),
            Collection::new(
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
}
