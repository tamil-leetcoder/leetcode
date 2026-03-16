# Valid Anagram

# Brute force
# def isAnagram(s, t):
#     return sorted(s) == sorted(t) # same -> true

# Optimal
def isAnagram(s,t):
    if len(s) != len(t):
        return False
    
    count = {}
    
    for char in s: # Increment
        count[char] = count.get(char, 0) + 1
        
    for char in t: # decrement
        count[char] = count.get(char, 0) - 1
        if count[char] < 0: # extra charcter
            return False
        
    return True

# Test 1
s = "east"
t = "seat"
print(isAnagram(s,t))