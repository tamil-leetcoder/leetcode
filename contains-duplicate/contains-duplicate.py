# 217 - Contains Duplicate

#Brute force
# def containsDuplicate(nums):
#     for i in range(len(nums)):
#         for j in range(i+1,len(nums)):
#             if nums[i] == nums[j]:
#                 return True   
#     return False

# Optimal
def containsDuplicate(nums):
    seen = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)
    return False

# One liner
# def containsDuplicate(nums):
#     return len(nums) != len(set(nums))

nums = [1,2,3,1]
print(containsDuplicate(nums))