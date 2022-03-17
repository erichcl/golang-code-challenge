import React from 'react';
import { Beer } from './Beer';


export const Beers = ({products}) => {
    return <>
        <h2>Beers</h2>
        <table>
          <thead>
            <tr>
              <th align="left">Product</th>
              <th align="left">Temperature</th>
              <th align="left">Status</th>
            </tr>
          </thead>
          <tbody>
            {products.map((product) => (
              <Beer product={product} key={product.id}/>
            ))}
          </tbody>
        </table>
    </>
}
