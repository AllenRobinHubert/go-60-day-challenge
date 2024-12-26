

const nums = [2, 7, 11, 15]


function twosum(arr, target){
    let hash = []

    for (let index = 0; index < arr.length; index++) {
        const complement = target - arr[index];
        // console.log(complement)
        hash.push({seen: arr[index], complement})
        hash.map((item) => {
            if ( item.complement == arr[index] ){
                console.log(item.complement, item.seen)
                return
            }
        })
    }



}

twosum(nums, 9)