async function exampleFetch() {
  const response = await fetch("https://data.messari.io/api/v1/news");
  const json = await response.json();
  console.log(json);
}

exampleFetch();
