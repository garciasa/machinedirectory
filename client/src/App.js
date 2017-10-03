import React from 'react';
import 'whatwg-fetch';

import Search from './Search';
import ListResults from './ListResults';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
    };
    this.handleSearch = this.handleSearch.bind(this);
  }

  handleSearch(tags) {
    if (tags.length < 3) {
      this.setState({ items: [] });
      return;
    }
    fetch(`http://localhost:3000/items/${tags}`)
      .then((response) => {
        return response.json();
      })
      .then((json) => {
        this.setState({ items: json.response });
      })
      .catch((err) => {
        console.log(err);
      });
  }

  render() {
    const { items } = this.state;
    return (
      <div>
        <Search onChange={this.handleSearch} />
        <br />
        <ListResults results={items} />
      </div>
    );
  }
}

export default App;
