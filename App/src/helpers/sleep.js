function sleep (time) {
  const resolution = new Promise((resolve) => setTimeout(resolve, time));
  return resolution;
}

export default sleep;
