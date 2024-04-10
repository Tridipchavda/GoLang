// TrainPage.js

import React, { useEffect, useState } from "react";
import { Button } from "./button";


const TrainPage = () => {
  // States to manage the Train record from DB , count of Paging Skip and search Query from input field
  const [searchQuery, setSearchQuery] = useState("");
  const [trains, setTrains] = useState([{}]);
  const [count, setCount] = useState(0);

  // Perform the Query to Server when the count of Paging Skip changes
  useEffect(() => {
    // Fetch the records according to the Skip
    fetch("http://localhost:8080/" + (count + 1))
      .then((response) => response.json())
      .then((responseData) => {
        // If the response data is comping set to the State Trains else Empty the Train Record
        if (responseData != "null") {
          setTrains(responseData);
        } else {
          setTrains([]);
        }
      })
      // Handling Error and console it
      .catch((error) => console.log(error));
  }, [count]);

  // Function to increament count
  const increment = () => {
    setCount(count + 1);
  };
  // Function to decreament count
  const decrement = () => {
    if (count != 0) {
      setCount(count - 1);
    } else {
      alert("Negative Paging Not Possible");
    }
  };
  // Handle the search from the database Query
  const findTrains = (e) => {
    // Change the input field and state value
    setSearchQuery(e.target.value);
    // If search field is empty then show the paging Data Accordingly
    if (e.target.value == "") {
      console.log(e.target.value);
      // If count = 0 set count to count + 1 else set count to count - 1
      if (count == 0) {
        setCount(count + 1);
      } else {
        setCount(count - 1);
      }
    } else {
      // Fetch the result for finding the search result at server for all Records in Database
      fetch("http://localhost:8080/find/" + e.target.value)
        .then((response) => response.json())
        .then((responseData) => setTrains(responseData))
        .catch((error) => console.log(error));
    }
  };

  return (
    <div className="bg-gray-100 min-h-screen py-8">
      <div className="container mx-auto px-4">
        <h1 className="text-3xl font-semibold text-center mb-6">
            Train Routes
        </h1>
        <div className="mb-8">
          <input
            type="text"
            placeholder="Search Train"
            className="w-6/12 border border-gray-300 rounded-md py-2 px-3 focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={searchQuery}
            onChange={(e) => findTrains(e)}
          />
        </div>
        <div className="overflow-x-auto">
          <table className="w-full bg-white border-collapse border border-gray-300 rounded-md">
            <thead className="bg-blue-500 text-white">
              <tr key="head">
                <th className="px-4 py-2">#</th>
                <th className="px-4 py-2">Train No</th>
                <th className="px-4 py-2">Train Name</th>
                <th className="px-4 py-2">Starts</th>
                <th className="px-4 py-2">Ends</th>
              </tr>
            </thead>
            <tbody>
              {trains != null ? (
                trains.map((train,index) => (
                  <tr
                    key={index}
                    className={
                      (train.index % 2 === 0 ? "bg-gray-50" : "bg-white") +
                      " hover:bg-gray-100"
                    }
                  >
                    <td className="px-4 py-2">{train.index}</td>
                    <td className="px-4 py-2">{train.train_no}</td>
                    <td className="px-4 py-2">{train.train_name}</td>
                    <td className="px-4 py-2">{train.starts}</td>
                    <td className="px-4 py-2">{train.ends}</td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td colSpan="5" className="m-3 p-3">
                    {" "}
                    No Data Found{" "}
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </div>
      <div className="flex justify-center my-8">
        <Button onClick={decrement}>Prev</Button>
        <span className="mt-2 text-md">{count}</span>
        <Button onClick={increment}>Next</Button>
      </div>
    </div>
  );
};

export default TrainPage;
