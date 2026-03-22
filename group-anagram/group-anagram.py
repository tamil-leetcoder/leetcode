# 49. Group Anagram

# Brute Force -> ["eat","tea","tan","ate","nat","bat"]

# def groupAnagram(strs):
#     result = [] # [[eat, tea, ate], [tan, nat], [bat]]
#     visited = [False] * len(strs) #[False, True, False, True, True, False]   
#     for i in range(len(strs)):
#         if visited[i]:
#             continue
#         group = [strs[i]] #bat
#         for j in range(i+1, len(strs)):
#             if sorted(strs[i]) == sorted(strs[j]):
#                 group.append(strs[j])
#                 visited[j] = True
#         result.append(group)   
#     return result


# Sorting Key 
# def groupAnagrams(strs): # O(n*klogk)
#     groups = {} 
#     for word in strs:
#         key = ''.join(sorted(word)) # klogk
#         if key not in groups:
#             groups[key] = []
#         groups[key].append(word)      
#     return list(groups.values())

# Character Count Array
def groupAnagrams(strs): # 
    groups = {}
    
    for word in strs:
        count = [0] * 26
        
        for char in word:
            count[ord(char) - ord('a')] += 1
        
        key = tuple(count) 
        if key not in groups:
            groups[key] = []
        groups[key].append(word)
        
    return list(groups.values())


strs = ["eat","tea","tan","ate","nat","bat"]
print(groupAnagrams(strs))