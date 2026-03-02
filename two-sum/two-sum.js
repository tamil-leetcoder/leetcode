/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */

/**
 * seen[complement] -> 0 -> false
 */
var twoSum = function(nums, target) {
    const seen = {};

    for(let i=0; i< nums.length; i++) {
        const complement = target - nums[i];

        if(seen[complement] !== undefined) {
            return [seen[complement], i];
        }

        seen[nums[i]] = i;
    }
};

const nums = [2,7,11,19];
const target = 9;

console.log(twoSum(nums, target));