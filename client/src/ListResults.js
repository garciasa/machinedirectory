import React from 'react';
import PropTypes from 'prop-types';

function ListResults({ results }) {
  const renderList = () => {
    return (
      <ul>
        {results.map((i) => {
          return (
            <li key={i.ip}>
              {i.domainname} - {i.ip} - {i.tags}
            </li>
          );
        })}
      </ul>
    );
  };
  return renderList();
}

ListResults.propTypes = {
  results: PropTypes.arrayOf(PropTypes.object).isRequired,
};

export default ListResults;
