function dfs(curr, cityCanLeadToZero, mapConnections, visited, result) {
  if (visited[curr]) {
    return;
  }
  visited[curr] = true;
  const connections = mapConnections[curr];
  for (const connection of connections) {
    const firstCity = connection[0];
    const secondCity = connection[1];
    if (cityCanLeadToZero[secondCity]) {
      cityCanLeadToZero[firstCity] = true;
      dfs(firstCity, cityCanLeadToZero, mapConnections, visited, result);
    } else if (cityCanLeadToZero[firstCity]) {
      result[0]++;
      cityCanLeadToZero[secondCity] = true;
      dfs(secondCity, cityCanLeadToZero, mapConnections, visited, result);
    }
  }
}

function addToMapConnections(city, mapConnections, connection) {
  if (mapConnections[city]) {
    mapConnections[city].push(connection);
  } else {
    mapConnections[city] = [connection];
  }
}

/**
 * @param {number} n
 * @param {number[][]} connections
 * @return {number}
 */
var minReorder = function(n, connections) {
  const mapConnections = Array(n);
  for (let i = 0; i < connections.length; i++) {
    const connection = connections[i];
    addToMapConnections(connection[0], mapConnections, connection);
    addToMapConnections(connection[1], mapConnections, connection);
  }

  const cityCanLeadToZero = Array(n).fill(false);
  cityCanLeadToZero[0] = true;
  const visited = {};
  const result = [0]; // act like pointer because js doesnt have one
  for (let i = 0; i < n; i++) {
    dfs(i, cityCanLeadToZero, mapConnections, visited, result);
  }
  return result[0];
};
