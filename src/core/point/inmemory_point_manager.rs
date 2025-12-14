use crate::proto::vectordb::Point;
use std::collections::HashMap;
use super::point_manager::PointManager;

#[derive(Debug)]
pub struct InMemoryPointManager {
    // database_name -> collection_name -> point_id -> Point
    points: HashMap<String, HashMap<String, HashMap<String, Point>>>,
}

impl InMemoryPointManager {
    /// Calculate Euclidean distance between two vectors
    fn euclidean_distance(vec1: &[f32], vec2: &[f32]) -> f32 {
        if vec1.len() != vec2.len() {
            return f32::INFINITY;
        }
        vec1.iter()
            .zip(vec2.iter())
            .map(|(a, b)| (a - b).powi(2))
            .sum::<f32>()
            .sqrt()
    }
    
    fn create_point(
        id: String,
        collection_name: String,
        database_name: String,
        vector: Vec<f32>,
    ) -> Point {
        let now = chrono::Utc::now().timestamp();
        Point {
            id,
            collection_name,
            database_name,
            vector,
            metadata: std::collections::HashMap::new(),
            created_at: now,
            updated_at: now,
        }
    }
}

impl PointManager for InMemoryPointManager {
    fn new() -> Self {
        let mut manager = Self {
            points: HashMap::new(),
        };
        
        // Add sample points for testing
        let sample_points = vec![
            Self::create_point(
                "point_1".to_string(),
                "products".to_string(),
                "default".to_string(),
                vec![0.1, 0.2, 0.3, 0.4],
            ),
            Self::create_point(
                "point_2".to_string(),
                "products".to_string(),
                "default".to_string(),
                vec![0.15, 0.25, 0.35, 0.45],
            ),
            Self::create_point(
                "point_3".to_string(),
                "products".to_string(),
                "default".to_string(),
                vec![0.9, 0.8, 0.7, 0.6],
            ),
        ];
        
        let mut collection_points = HashMap::new();
        for point in sample_points {
            collection_points.insert(point.id.clone(), point);
        }
        
        let mut db_collections = HashMap::new();
        db_collections.insert("products".to_string(), collection_points);
        
        manager.points.insert("default".to_string(), db_collections);
        
        manager
    }
    
    fn get_point(&self, database_name: &str, collection_name: &str, point_id: &str) -> Option<&Point> {
        self.points
            .get(database_name)
            .and_then(|db| db.get(collection_name))
            .and_then(|col| col.get(point_id))
    }
    
    fn search_nearest(
        &self,
        database_name: &str,
        collection_name: &str,
        query_vector: &[f32],
        limit: usize,
    ) -> Vec<(Point, f32)> {
        let mut results = Vec::new();
        
        if let Some(db) = self.points.get(database_name) {
            if let Some(collection) = db.get(collection_name) {
                for point in collection.values() {
                    let distance = Self::euclidean_distance(&point.vector, query_vector);
                    results.push((point.clone(), distance));
                }
            }
        }
        
        // Sort by distance (ascending) and limit results
        results.sort_by(|a, b| a.1.partial_cmp(&b.1).unwrap_or(std::cmp::Ordering::Equal));
        results.truncate(limit);
        
        results
    }
    
    fn add_point(&mut self, point: Point) -> Result<(), String> {
        // Validate that the database exists
        if !self.points.contains_key(&point.database_name) {
            return Err(format!("Database '{}' does not exist", point.database_name));
        }
        
        // Validate that the collection exists
        if !self.points[&point.database_name].contains_key(&point.collection_name) {
            return Err(format!(
                "Collection '{}' does not exist in database '{}'",
                point.collection_name, point.database_name
            ));
        }
        
        // Check if point already exists
        if self.points[&point.database_name][&point.collection_name].contains_key(&point.id) {
            return Err(format!(
                "Point '{}' already exists in collection '{}' in database '{}'",
                point.id, point.collection_name, point.database_name
            ));
        }
        
        // Add the point
        self.points
            .get_mut(&point.database_name)
            .unwrap()
            .get_mut(&point.collection_name)
            .unwrap()
            .insert(point.id.clone(), point);
        
        Ok(())
    }
}
