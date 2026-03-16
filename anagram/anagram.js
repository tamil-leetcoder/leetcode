// Valid Anagram

// Brute Force
// function isAnagram(s, t) {
//     return s.split('').sort().join('') === t.split('').sort().join('') // true
// }

// Optimal Solution
function isAnagram (s, t) {
    if (s.length !== t.length) {
        return false
    }

    const count = {}

    for (let char of s) {
        // increment
        count[char] = (count[char] || 0) + 1;
    }

    for (let char of t) {
        // decrement
        count[char] = (count[char] || 0) - 1;
        if(count[char] < 0) {
            return false
        }
    }

    return true;
}

const s = "east";
const t = "seat";

console.log(isAnagram(s, t))