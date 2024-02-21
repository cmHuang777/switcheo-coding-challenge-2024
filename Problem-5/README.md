### What Does It Mean by Breaking Consensus?
Breaking consensus in a blockchain context refers to making a change to the protocol or network rules that results in a lack of agreement among nodes in the network. Consensus is the shared agreement among nodes on the state of the blockchain, including valid transactions and the order in which they are added to the chain. A consensus-breaking change disrupts this agreement, leading to potential network forks and inconsistencies in the blockchain's state.

### Why Your Change Would Break the Consensus
Code changes: 
```
message Post {
  
  string title = 1; 
  string body = 2; 
  string creator = 3; 
  uint64 id = 4; 
}
```
To 
```
message Post {
  
  string title = 2; 
  string body = 1; 
  string creator = 3; 
  uint64 id = 4; 
}
```
This changes the protocol on post directly and should break the consensus with existing nodes. This causes a disagreement on how post should be like. 