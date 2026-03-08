// Brute force
// const containsDuplicate = function (nums) {
//     for (let i=0; i<nums.length;i++) {
//         for(let j=i+1;j<nums.length;j++) {
//             if(nums[i] === nums[j]) {
//                 return true;
//             }
//         }
//     }
//     return false;
// }

// optimal
const containsDuplicate = function (nums) {
    const seen = new Set()

    for (let num of nums) {
        if(seen.has(num)) {
            return true
        }
        seen.add(num);
    }

    return false;
}

// one liner
const oneLiner = function(nums) {
    return new Set(nums).size !== nums.length
}

const nums = [1,2,3,1]

console.log(containsDuplicate(nums))