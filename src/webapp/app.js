import React, { useState } from 'react';
import Graph from 'react-graph-vis';
import './app.css';
import ReactImage from './react.png';

const App = () => {
  const [input, setInput] = useState("A B\nB C\nC A\nB D\nD E\nE F\nF E");
  const [graph, setGraph] = useState(null);
  const [runtime, setRuntime] = useState(null); // State for storing the runtime

  const colors = [
    "red",
    "blue",
    "green",
    "yellow",
    "purple",
    "orange",
    "gray",
    "cyan",
    "magenta",
    "lime",
    "pink",
    "teal",
    "indigo",
    "maroon",
    "olive",
    "navy",
    "silver",
    "aqua",
    "fuchsia",
    "limegreen",
    "violet",
    "gold",
    "coral",
    "steelblue",
  ];
  const handleInputChange = (event) => {
    setInput(event.target.value);
  };


  const handleSubmit = () => {
    const startTime = performance.now(); // Start time of the request

    fetch('/api/scc', {
      method: 'POST',
      body: JSON.stringify({ input }),
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((res) => res.json())
      .then((response) => {
        const endTime = performance.now(); // End time of the request
        const runtime = endTime - startTime; // Calculate the runtime in milliseconds
        setRuntime(runtime.toFixed(2)); // Set the runtime state

        const nodes = [];
        const edges = [];

        // Process nodes
        for (let i = 0; i < response.nodes.length; i++) {
          const node = response.nodes[i];
          const component = node.component;
          const label = `Node ${node.id}`;

          nodes.push({
            id: node.id,
            label,
            component,
            color: colors[component % colors.length],
          });
        }

        // Process edges
        for (let i = 0; i < response.edges.length; i++) {
          const edge = response.edges[i];
          const from = edge.from;
          const to = edge.to;
          const isBridge = edge.bridges;

          edges.push({
            from,
            to,
            bridges: isBridge,
            color: isBridge ? 'blue' : 'red',
          });
        }

        const graphData = {
          nodes,
          edges,
        };
        console.log(graphData)

        setGraph(graphData);
      })
      .catch((error) => console.error(error));
  };

  const options = {
    layout: {
      hierarchical: false,
    },
    edges: {
      width: 1,
      arrows: {
        to: { enabled: true, scaleFactor: 0.5 },
      },
    },
  };

  const events = {};
  const handleFileChange = (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.onload = (e) => {
      const fileContent = e.target.result;
      setInput(fileContent);
    };
    reader.readAsText(file);
  };

  return (
    <div className="container">
      <h1 className="title">Graph Visualization</h1>
      <div className="input-section">
        <label htmlFor="input" className="input-label">
          Graph Input:
        </label>
        <textarea
          id="input"
          value={input}
          onChange={handleInputChange}
          className="input-textarea"
        />
        <input
          type="file"
          accept=".txt"
          onChange={handleFileChange}
          className="input-file"
        />
        <button onClick={handleSubmit} className="submit-button">
          Submit
        </button>
      </div>
      <div className="runtime">
        {runtime && <p>Runtime: {runtime} ms</p>}
      </div>
      <div className="graph-container">
        {graph && <Graph graph={graph} options={options} events={events} />}
      </div>
    </div>
  );

};

export default App;
