type LocationResponse = {
  Feature: Array<{
    Geometry: {
      Type: string;
      Coordinates: string;
    };
    Property: {
      Address: string;
    };
  }>;
};
