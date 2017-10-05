import React from 'react';
import PropTypes from 'prop-types';

function ListResults({ results }) {
  const renderList = () => {
    return (
      <ul>
        {results.map((i) => {
          return (
            <li key={i.ip}>
              <div className="domainname">{i.domainname}</div>
              <div className="ip">{i.ip}</div>
              <div className="tags">{i.tags}</div>
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
