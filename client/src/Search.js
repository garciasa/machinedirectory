import React from 'react';
import PropTypes from 'prop-types';


function Search({ onChange }) {
  function handleChange(e) {
    e.preventDefault();
    onChange(e.target.value);
  }

  return <input type="text" placeholder="search" onChange={handleChange} />;
}

Search.propTypes = {
  onChange: PropTypes.func.isRequired,
};

export default Search;
