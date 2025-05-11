// Helper function to measure execution time
function measureTime(fn) {
    const start = process.hrtime.bigint();
    fn();
    const end = process.hrtime.bigint();
    return Number(end - start) / 1000000; // Convert to milliseconds
}

// Create a large array of random numbers
const DATA_SIZE = 10000000;
const data = new Array(DATA_SIZE);
for (let i = 0; i < DATA_SIZE; i++) {
    data[i] = Math.random() * 256;
}

// Sort the data for predictable branches
const sortedData = [...data].sort((a, b) => a - b);

// Test 1: Unsorted data (unpredictable branches)
console.log("Testing with UNSORTED data:");
const unsortedTime = measureTime(() => {
    let sum = 0;
    for (let i = 0; i < DATA_SIZE; i++) {
        if (data[i] >= 128) {  // Unpredictable branch
            sum += data[i];
        }
    }
    console.log("Sum (unsorted):", sum);
});
console.log("Time taken (unsorted):", unsortedTime.toFixed(2), "ms\n");

// Test 2: Sorted data (predictable branches)
console.log("Testing with SORTED data:");
const sortedTime = measureTime(() => {
    let sum = 0;
    for (let i = 0; i < DATA_SIZE; i++) {
        if (sortedData[i] >= 128) {  // Predictable branch (all false, then all true)
            sum += sortedData[i];
        }
    }
    console.log("Sum (sorted):", sum);
});
console.log("Time taken (sorted):", sortedTime.toFixed(2), "ms\n");

// Test 3: Without any branches (for comparison)
console.log("Testing WITHOUT branches:");
const noBranchTime = measureTime(() => {
    let sum = 0;
    for (let i = 0; i < DATA_SIZE; i++) {
        // Use bitwise operations to avoid branching
        sum += data[i] * ((data[i] >= 128) | 0);
    }
    console.log("Sum (no branch):", sum);
});
console.log("Time taken (no branch):", noBranchTime.toFixed(2), "ms\n");

// Compare results
console.log("Performance comparison:");
console.log(`Sorted is ${(unsortedTime / sortedTime).toFixed(2)}x faster than unsorted`);
console.log(`No-branch is ${(unsortedTime / noBranchTime).toFixed(2)}x relative to unsorted`);
