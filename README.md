[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/Dt3ukIt2)

# Parallel Mergesort

Implement a parallel version of mergesort (both the original recursive and the
iterative in-place version from a previous exercise are fine). You may use any
parallelization framework or method.

I have not provided any test code, but you can base yours on test code from
other exercises. Your tests must check the correctness of the result of running
the function and run automatically when you commit through a GitHub action.

## Runtime Analysis

What is the span of the parallel program, in terms of worst-case $\Theta$? Hint:
It may help to consider the DAG of the parallel program.

**FINAL**

With the sequential version of mergesort, we have a time complexity of $\Theta(n \cdot \log ⁡n)$. However, when analyzing the span of this parallel version, we must consider the critical path's length. In my implementation, the critical path is essentially the merge step. Each time I partition the array and delegate the task to another process, a sequential operation of merging is required. This means a time complexity of $\Theta(n)$ during the merge process.

Since we do not have an infinite number of processors, we are limited by the number of cores we can use in practice. An approach that would work given an infinite number of cores would be to separate each piece of work into its own core and repeat this process until each core works on one piece of data. Then, the final merge would occur.

However, since my code is based in reality, I opted to cease further parallel work after reaching a number determined by the user's core count. The overall analysis should, in theory, be the same as traditional mergesort, as we are performing the same amount of work. The merge step is $\Theta(⁡n)$ and the partitioning is $\Theta( \log ⁡n)$.

Thus, in practice, we are not doing any more or less work. In the real world, we manage threads and other operational overheads which asymptotic analysis does not account for. Therefore, the worst-case analysis remains the same between both the parallel and sequential versions. The worst-case time complexity of my parallel mergesort is $\Theta(n \cdot \log ⁡n)$.
