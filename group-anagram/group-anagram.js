// 49. Group Anagram

// Brute Force Approach
// function groupAnagrams(strs) {
//     const result = [];
//     const visited = new Array(strs.length).fill(false);

//     for (let i=0; i<strs.length; i++) {
//         if (visited[i]) continue

//         const group = [strs[i]]
//         for(let j=i+1; j<strs.length; j++) {
//             if(strs[i].split('').sort().join('') === 
//                 strs[j].split('').sort().join('')) {
//                     group.push(strs[j]);
//                     visited[j] = true
//                 }
//         }

//         result.push(group)
//     }

//     return result
// }

// Sorting Key Approach
// function groupAnagrams(strs) {
//     const groups = {}; // Time complexity: O(n*klogk)

//     for (let word of strs) {
//         const key = word.split('').sort().join(''); // klogk
//         groups[key] = groups[key] || [];
//         groups[key].push(word)
//     }

//     return Object.values(groups)
// }

// Character Count Array - optimal
function groupAnagrams(strs) {
    const groups = {} // Time complexity: O(n*k)

    for (let word of strs) {
        const count = new Array(26).fill(0);

        for (let char of word) {
            count[char.charCodeAt(0) - 97] += 1;
        }

        const key = count.join(',');

        if(!groups[key]) {
            groups[key] = []
        }

        groups[key].push(word)
    }

    return Object.values(groups)
}

const strs = ["eat","tea","tan","ate","nat","bat"]
console.log(groupAnagrams(strs))



