import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import * as Api from './Api';

class Add extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      ip: '',
      domainname: '',
      tags: '',
    };
    this.cancel = this.cancel.bind(this);
    this.input = this.input.bind(this);
    this.save = this.save.bind(this);
  }
  cancel(evt) {
    evt.preventDefault();
    this.props.history.push('/');
  }

  input(evt) {
    evt.preventDefault();
    const target = evt.target;
    const { value, name } = target;
    this.setState({
      [name]: value,
    });
  }
  save(evt) {
    evt.preventDefault();
    Api.Save(this.state)
      .then((result) => {
        console.log(result);
        this.props.history.push('/');
      });
  }

  render() {
    return (
      <div>
        <Link to="/">Back</Link>
        <div className="form-add">
          <input type="text" name="ip" onChange={this.input} placeholder="IP" />
          <input
            type="text"
            name="domainname"
            onChange={this.input}
            placeholder="Domain name"
          />
          <input
            type="text"
            name="tags"
            placeholder="tags separated by commas"
            onChange={this.input}
          />
        </div>
        <div className="control-add-buttons">
          <button onClick={this.save}>Save</button>
          <button onClick={this.cancel}>Cancel</button>
        </div>
      </div>
    );
  }
}

Add.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func.isRequired,
  }).isRequired,
};

export default Add;
