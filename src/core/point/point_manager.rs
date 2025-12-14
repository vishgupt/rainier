use crate::proto::vectordb::Point;

/// Defines core point/vector management operations.
/// Implement this trait for different storage backends.
pub trait PointManager: Send + Sync + 'static {
    /// Creates a new instance of the point manager.
    fn new() -> Self where Self: Sized;
    
    /// Retrieves a specific point by ID.
    fn get_point(&self, database_name: &str, collection_name: &str, point_id: &str) -> Option<&Point>;
    
    /// Finds nearest neighbors to a given vector.
    /// Returns up to `limit` nearest points sorted by distance (closest first).
    fn search_nearest(
        &self,
        database_name: &str,
        collection_name: &str,
        query_vector: &[f32],
        limit: usize,
    ) -> Vec<(Point, f32)>; // Returns (Point, distance) tuples
    
    /// Adds a new point to a collection.
    fn add_point(&mut self, point: Point) -> Result<(), String>;
}
