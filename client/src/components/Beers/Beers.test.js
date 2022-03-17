import React from "react";
import { render, unmountComponentAtNode } from "react-dom";
import { screen, waitFor } from '@testing-library/react';
import { toHaveAttribute } from '@testing-library/jest-dom/matchers'
import '@testing-library/jest-dom/extend-expect'
import { act } from "react-dom/test-utils";
import { Beers } from "./Beers"

let container = null;
let tbody = null;
beforeEach(() => {
  container = document.createElement("div");
  document.body.appendChild(container);
});

afterEach(() => {
  // cleanup on exiting
  unmountComponentAtNode(container);
  container.remove();
  container = null;
});

it("renders empty without crashing", async () => {
  const fakeProductList = [];
  jest.spyOn(global, "fetch").mockImplementation(() =>
    Promise.resolve({
      json: () => Promise.resolve(fakeProduct)
    })
  );

  // Use the asynchronous version of act to apply resolved promises
  await act(async () => {
    render(<Beers products={fakeProductList} />, container);
  });
  // remove the mock to ensure tests are completely isolated
  global.fetch.mockRestore();
});


it("renders data without crashing", async () => {
  const fakeProductList = [
    {id: 12345,
    name: "Johnny Beer",
    temperature: "6",
    temperatureStatus: "Just Fine"}
  ];
  jest.spyOn(global, "fetch").mockImplementation(() =>
    Promise.resolve({
      json: () => Promise.resolve(fakeProduct)
    })
  );

  // Use the asynchronous version of act to apply resolved promises
  await act(async () => {
    render(<Beers products={fakeProductList} />, container);
  });

    expect(screen.getByTestId("td-name").textContent).toBe(fakeProductList[0].name);
    expect(screen.getByTestId("td-temperature").textContent).toBe(fakeProductList[0].temperature.toString());
    expect(screen.getByTestId("td-temperatureStatus").textContent).toBe(fakeProductList[0].temperatureStatus);

  // remove the mock to ensure tests are completely isolated
  global.fetch.mockRestore();
});