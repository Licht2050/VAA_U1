# VAA_Uebung1


- The following program start a cluster with master and worker nodes
- Each node first randomly selects neighbors and sends information about them to the master node.
- The master node allows the user to print all nodes with neighbors.
- The master node also allows the user to start sending rumors. and each node sends the rumors to its neighbors if it receives it for the first time.
- The user can also print information about neighboring nodes as a directed graph. it can be printed to stdr, write it to a ".dot" file, or save it as a "png" file.
- In addition, the Master node allows the user to create an undirected graph.

## There is two way to start the program:

1. It can provide a ".txt" file path as a parameter, which includes information about the cluster master and the worker node. The following example shows the format, such as the information tree to be recorded.

    **Master: 127.0.0.1:8793**\
    **node02: 127.0.0.1:8794**
    
    The program takes as the first parameter a "file" string, and the second parameter the path of the file as follows: 

    **./start file Nodes.txt**
    
    __A "Nodes.txt" file already exists in the root folder.__


2. could automatically start any node number. The program accepts two parameters, the first an "auto" string and the second an integer, which represents the number of nodes. The following example starts a cluster with 10 worker nodes:

    **./start auto 10**


The program start master and worker nodes in a new tabs. each node shows the cycle of random neighbor choice. the neighbor selection process begins each time a new node has successfully joined the cluster.

A menu appears on the master node 10 seconds after the cluster starts, which can help you get more information about the cluster.
