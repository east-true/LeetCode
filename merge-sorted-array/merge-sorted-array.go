func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i := m; i < n+m; i++ {
        nums1[i] = nums2[i-m]
    }
    
    sort.Ints(nums1)
}