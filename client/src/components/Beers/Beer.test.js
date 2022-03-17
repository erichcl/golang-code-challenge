import React from "react";
import { render, unmountComponentAtNode } from "react-dom";
import { screen, waitFor } from '@testing-library/react';
import { toHaveAttribute } from '@testing-library/jest-dom/matchers'
import '@testing-library/jest-dom/extend-expect'
import { act } from "react-dom/test-utils";
import { Beer } from "./Beer"

let container = null;
let tbody = null;
beforeEach(() => {
  container = document.createElement("table");
  tbody = document.createElement("tbody");
  container.appendChild(tbody);
  document.body.appendChild(container);
});

afterEach(() => {
  // cleanup on exiting
  unmountComponentAtNode(tbody);
  tbody.remove();
  tbody = null;
});

it("renders product data", async () => {
  const fakeProduct = {
    id: 12345,
    name: "Johnny Beer",
    temperature: "6",
    temperatureStatus: "Just Fine"
  };
  jest.spyOn(global, "fetch").mockImplementation(() =>
    Promise.resolve({
      json: () => Promise.resolve(fakeProduct)
    })
  );

  // Use the asynchronous version of act to apply resolved promises
  await act(async () => {
    render(<Beer product={fakeProduct} />, tbody);
  });

    expect(screen.getByTestId("td-name").textContent).toBe(fakeProduct.name);
    expect(screen.getByTestId("td-temperature").textContent).toBe(fakeProduct.temperature.toString());
    expect(screen.getByTestId("td-temperatureStatus").textContent).toBe(fakeProduct.temperatureStatus);

  // remove the mock to ensure tests are completely isolated
  global.fetch.mockRestore();
});