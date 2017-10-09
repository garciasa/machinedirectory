import React from 'react';

import Search from './Search';
import ListResults from './ListResults';
import * as Api from './Api';

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
    Api.getItems(tags)
      .then((data) => {
        this.setState({ items: data.response });
      });
  }

  render() {
    const { items } = this.state;
    return (
      <div id="wrapper">
        <Search onChange={this.handleSearch} />
        <ListResults results={items} />
      </div>
    );
  }
}

export default App;
