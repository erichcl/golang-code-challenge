import React, { Component } from 'react';
import { Beers } from './components/Beers/Beers';

export class App extends Component {
  static displayName = App.name;

  constructor(props) {
    super(props);
    this.state = { products: [], loading: true };
  }

  componentDidMount() {
    setInterval(() => {
      this.fetchProducts();
    }, 5000);
  }

  async fetchProducts() {
    const response = await fetch('api/products');
    const data = await response.json();
    this.setState({ products: data, loading: false });
  }

  render() {
    let contents = this.state.loading
      ? (
        <p><em>Loading...</em></p>
      ) : (
      <>
        <Beers products={this.state.products}/>
      </>
    );

    return (
      <>
        <header>
          <h1>SensorTech</h1>
        </header>
        <main>
          {contents}
        </main>
      </>
    );
  }
}
