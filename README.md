# request-completion-time
**Efficient Horizontal Scaling Time in Clustered Computing Environments**

### Paper Information
- **Author(s):** SaiKrishna Mylavarapu
- **Published In:** Advanced International Journal of Multidisciplinary Research (AIJMR)
- **Publication Date:** Aug, 2023
- **ISSN:** E-ISSN: 2584-0487
- **DOI:** https://doi.org/10.62127/aijmr.2023.v01i01.1381
- **Impact Factor:** 9.11

### Abstract
Efficient replica synchronization is essential for maintaining consistency, fault tolerance, and high availability in multi node distributed architectures. Existing synchronization mechanisms experience increasing delays as cluster size and workload intensity grow, resulting in stale replica states, inconsistent processing, and reduced throughput. The study analyzes synchronization behavior across varying cluster sizes and identifies the limitations of static synchronization strategies. It introduces an enhanced synchronization management approach that continuously monitors synchronization activity, workload distribution, communication delay, and node resource utilization. The proposed mechanism dynamically adjusts synchronization operations to reduce replica delays, improve state consistency, and enhance overall synchronization efficiency in distributed environments.

**Core Technical Contributions**
- **Adaptive Replica Synchronization Framework:**
Designed an adaptive synchronization framework that dynamically regulates replica update propagation based on runtime synchronization delay, workload distribution, communication activity, and node resource utilization, improving synchronization efficiency across distributed architectures.
- **Runtime Synchronization Monitoring Mechanism:**
Implemented continuous monitoring of synchronization delay, communication behavior, replica health, and resource utilization to detect synchronization bottlenecks and maintain balanced replica coordination during runtime.
- **Adaptive Propagation Control Model:**
Developed a dynamic propagation mechanism that prioritizes replicas experiencing higher synchronization backlog, reducing propagation delay, synchronization backlog accumulation, and inconsistent replica state visibility.
- **Concurrent Multi-Replica Synchronization Simulator:**
Implemented a Go-based concurrent synchronization model using Goroutines and WaitGroups to simulate parallel replica synchronization, adaptive decision making, and runtime monitoring across distributed nodes.
- **Scalability Analysis Across Cluster Sizes:**
Evaluated synchronization performance across clusters containing 3, 5, 7, 9, and 11 nodes, demonstrating that adaptive synchronization maintains lower synchronization delay and better scalability than conventional static synchronization approaches.

### Experimental Results (Summary)

  | Nodes | Static Synchronization (ms) | Adaptive Synchronization (ms) | Improvement (%) |
  |-------|-----------------------------| ------------------------------| ----------------|
  | 3     |  420                        | 290                           | 30.95           |
  | 5     |  560                        | 360                           | 35.71           |
  | 7     |  710                        | 430                           | 39.44           |
  | 9     |  860                        | 510                           | 40.70           |
  | 11    |  1020                       | 590                           | 42.16           |

### Citation
Efficient Horizontal Scaling Time in Clustered Computing Environments
* SaiKrishna Mylavarapu
* Advanced International Journal of Multidisciplinary Research (AIJMR)  
* ISSN E-ISSN: 2584-0487
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.
* Resources \
https://www.aijmr.com/ 
* Author Contact \
  **LinkedIn**: linkedin.com/in/saikrishna-mylavarapu-35479114 | **Email**: krishnamysap@gmail.com






