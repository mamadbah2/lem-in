# 🐜 Lem-in - Ant Colony Pathfinding Simulator

An advanced algorithmic project that simulates an ant colony finding optimal paths through a graph structure. This project focuses on developing advanced algorithm design skills, including pathfinding, graph traversal, and optimization techniques.

![Go](https://img.shields.io/badge/Go-100%25-00ADD8?logo=go&logoColor=white)

## 🎯 Project Overview

Lem-in is a graph traversal and optimization problem where ants must move from a starting room to an ending room through a network of tunnels in the minimum number of turns. The challenge lies in finding multiple non-overlapping paths and distributing ants optimally to minimize total movement time.

### **The Problem**
- **Input**: A colony file containing:
  - Number of ants
  - Rooms (nodes) with coordinates
  - Tunnels (edges) connecting rooms
  - Start and end rooms
- **Goal**: Move all ants from start to end in the minimum number of turns
- **Constraint**: Only one ant per room at a time (except start/end)

## ✨ Key Features

### 🔍 Pathfinding Algorithm
- **Multiple path detection** - Find all valid paths from start to end
- **Graph traversal** - Depth-first search implementation
- **Path validation** - Ensure no overlapping paths (except start/end)
- **Optimal path selection** - Choose the best combination of paths

### 📊 Optimization Algorithms
- **Ant distribution** - Intelligent distribution of ants across paths
- **Load balancing** - Minimize total movement time
- **Turn minimization** - Calculate optimal number of moves
- **Path length analysis** - Compare and select efficient routes

### 🚀 Movement Simulation
- **Turn-by-turn simulation** - Visualize ant movements
- **Collision avoidance** - Manage room occupancy
- **Parallel path execution** - Multiple ants moving simultaneously

## 🛠️ Technology Stack

### **Language**
![Go](https://img.shields.io/badge/Go_1.21.3-00ADD8?style=for-the-badge&logo=go&logoColor=white)

- **Pure Go Implementation** - No external dependencies
- **Standard Library Only** - Leveraging Go's built-in packages
- **Efficient Data Structures** - Optimized for performance

### **Core Technologies**
- Graph representation and traversal
- Recursive pathfinding algorithms
- Dynamic programming optimization
- File parsing and validation

## 🚀 Usage

```bash
# Clone the repository
git clone https://github.com/mamadbah2/lem-in.git
cd lem-in

# Run with example file
go run main.go

# Output displays:
# 1. Original input file
# 2. Ant movements turn by turn (L<ant_number>-<room>)
```

### Example Input (`repository/example.txt`)
```text
10
##start
start 1 6
0 4 8
##end
end 11 6
start-0
0-end
```

### Example Output
```text
L1-0 L2-0
L1-end L2-end L3-0 L4-0
L3-end L4-end L5-0 L6-0
...
```

## 📁 Project Structure

```
lem-in/
├── main.go               # Entry point and orchestration
├── colony/
│   ├── parser.go        # File parsing and validation
│   ├── pathfinder.go    # Pathfinding algorithm (DFS)
│   └── colony.go        # Ant distribution & movement logic
├── models/
│   └── models.go        # Data structures (File, Room, Link)
├── pkg/
│   └── utils.go         # Helper functions
└── repository/
    └── example.txt      # Test cases
```

## 🎓 Advanced Algorithmic Skills Acquired

### **Graph Theory & Algorithms**
✅ **Graph Representation** - Adjacency list implementation  
✅ **Depth-First Search (DFS)** - Recursive pathfinding  
✅ **Path Detection** - Finding all valid paths in a graph  
✅ **Graph Traversal** - Node visitation and backtracking  
✅ **Cycle Detection** - Avoiding infinite loops in paths

### **Optimization Techniques**
✅ **Dynamic Programming** - Optimal ant distribution  
✅ **Load Balancing** - Distributing work across multiple paths  
✅ **Greedy Algorithms** - Path selection strategies  
✅ **Complexity Analysis** - Time/space optimization  
✅ **Heuristic Design** - Approximating optimal solutions

### **Data Structures**
✅ **Graph Structures** - Nodes and edges representation  
✅ **Hash Maps** - Efficient lookups and tracking  
✅ **Arrays & Slices** - Dynamic data management  
✅ **Structs** - Complex data modeling

### **Algorithm Design**
✅ **Recursive Algorithms** - Backtracking pathfinding  
✅ **Iterative Algorithms** - Movement simulation  
✅ **State Management** - Tracking visited nodes  
✅ **Edge Case Handling** - Validation and error management

### **Problem Solving**
✅ **Problem Decomposition** - Breaking down complex problems  
✅ **Algorithm Selection** - Choosing appropriate approaches  
✅ **Optimization Strategy** - Improving performance  
✅ **Testing & Validation** - Ensuring correctness

### **Go Programming**
✅ **Pointers & References** - Memory management  
✅ **Struct Methods** - Object-oriented patterns  
✅ **File I/O** - Reading and parsing text files  
✅ **String Manipulation** - Parsing complex formats  
✅ **Error Handling** - Robust error management

## 🧮 Algorithm Complexity

- **Pathfinding**: O(V + E) where V = rooms, E = tunnels
- **Path Distribution**: O(n × p) where n = ants, p = paths
- **Movement Simulation**: O(t × a) where t = turns, a = ants

## 🔧 Key Technical Achievements

1. **Efficient Pathfinding** - DFS implementation with backtracking
2. **Optimal Distribution** - Smart ant allocation across multiple paths
3. **Movement Optimization** - Minimizing total number of turns
4. **Graph Parsing** - Robust file format validation
5. **Pure Go Solution** - No external dependencies

## 📚 Learning Outcomes

This project demonstrates mastery of:
- Advanced graph algorithms
- Optimization techniques
- Recursive problem solving
- Complex data structure manipulation
- Algorithm complexity analysis
- Go programming best practices

## 👥 Author

**[mamadbah2](https://github.com/mamadbah2)**

## 📝 License

Open source project - Available for educational purposes.

---

**Built with 🧠 at Zone01 Dakar - Advanced Algorithms Track**
