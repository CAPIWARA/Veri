async function save (id, coordinate) {
  const response = await fetch(`//veri-205319.firebaseio.com/coordinates/${id}.json`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(coordinate)
  });

  if (response.status !== 200)
    console.log({ response });
}

export { save };
