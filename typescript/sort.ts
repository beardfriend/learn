interface Idata {
  name: string;
  created_at: string;
}

const data: Idata[] = [
  {
    name: '김',
    created_at: '2022-10-10 22:00:32'
  },
  {
    name: '박',
    created_at: '2022-10-09 22:00:32'
  },
  {
    name: '장',
    created_at: '2022-10-10 22:00:33'
  },
  {
    name: '철',
    created_at: '2021-10-10 23:00:33'
  },
  {
    name: '철',
    created_at: '2023-10-10 23:00:33'
  }
];
// 첫 번쨰 방법

const first = data.sort((a, b) => {
  const keyA = new Date(a.created_at);
  const keyB = new Date(b.created_at);
  if (keyA < keyB) return 1;
  if (keyA > keyB) return -1;
  return 0;
});

console.log(first);

// 두 번째 방법

const { compare } = Intl.Collator('ko-KR');
const second = data.sort((a, b) => compare(a.created_at, b.created_at));
console.log(second);
