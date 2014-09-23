# GO Implementation of "Clustering by fast search and find of density peaks"

## Details
This clustering algorithm is recently featured by Science[1]. It uses 
local density and distances between points and their high-density neighbors. 

This new clustering method has the following advantages.
1. Unlike iteration based algorithms, this one identifies clusters in one round. 
2. Hyper-parameters are easy to deal with.
3. It handles density-based clustering or centroid-based clustering cases well.

However, the trade-off is that you have to compute and/or store pairwise distances. 
Although the theoretical asymptotic complexity could be lower than finding optimal clusters in k-means,
it could be slower in practice.

This implementation borrows the idea from [2] to make the hyperparameter selection even simpler. 

## Reference
1. http://www.sciencemag.org/content/344/6191/1492.abstract
2. http://rseghers.com/machine-learning/rodriguez-laoi-clustering/
