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
Horizontal scaling plays a critical role in maintaining responsiveness in modern clustered computing environments under dynamic workloads. Existing scaling mechanisms experience delays caused by orchestration overhead, resource contention, communication costs, and ineffective scheduling, resulting in increased request completion time. The study analyzes scaling behavior across multiple node configurations to identify the factors contributing to scaling latency. It proposes an advanced architectural model that improves resource allocation, accelerates node integration, and optimizes synchronization during scaling operations. The findings demonstrate that reducing scaling time enhances system responsiveness, operational stability, and scalability while providing a time-centric framework for evaluating clustered computing performance.

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

  | Nodes | Conventional Archtecture (ms) | Advanced Architecture (ms) | Improvement (%) |
  |-------|-------------------------------| ---------------------------| ----------------|
  | 3     |  1200                         | 480                        | 60.00           |
  | 5     |  1050                         | 390                        | 62.86           |
  | 7     |  980                          | 340                        | 65.31           |
  | 9     |  940                          | 310                        | 67.02           |
  | 11    |  910                          | 280                        | 69.23           |

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






