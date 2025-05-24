async function readArr(start: number, end: number, a: number[]) {
  let sum = 0;
  for (let i = start; i < end; i++) {
    sum += a[i];
  }
  console.log(`IGI ${start} ${end} ${sum}`);
  return sum;
}

async function main() {
  let a: number[];
  a = [];

  for (let i = 0; i < 100; i++) {
    a.push(i + 1);
  }

  let tasks: Promise<number>[];
  tasks = [];

  for (let i = 0; i < 100; i += 10) {
    tasks.push(readArr(i, i + 10, a));
  }

  let totalSum: number = 0;

  let result = await Promise.all(tasks);

  result.forEach((curr) => {
    totalSum += curr;
  });

  console.log(totalSum);
  return totalSum;
  //   for (let i = 0; i < 100; i += 10) {
  //     let currSum = await readArr(i, i + 10, a);
  //     totalSum += currSum;
  //   }

  //   console.log(totalSum);

  //   tasks.forEach((task, ind) => {
  //     task.then((currSum) => {
  //       totalSum += currSum;
  //     });

  //     if (ind == tasks.length) return totalSum;
  //   });
}

main().then((d) => {
  console.log(d);
});
