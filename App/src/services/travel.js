function toPath (coordinate) {
  return {
    lat: coordinate.latitude,
    lng: coordinate.longitude
  };
}

export { toPath };
