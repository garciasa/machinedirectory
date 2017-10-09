import 'whatwg-fetch';
import { API_URL } from './Config';

export const getItems = (tags) => {
  const result = new Promise((resolve) => {
    const url = `${API_URL}/items/${tags}`;
    fetch(url)
      .then((response) => {
        return response.json();
      })
      .then((json) => {
        resolve(json);
      })
      .catch((err) => {
        // TODO: collect erros in a single file
        console.log(err);
      });
  });
  return result;
};

export const Save = (data) => {
  const result = new Promise((resolve) => {
    const url = `${API_URL}/item/`;
    fetch(url,
      {
        method: 'POST',
        body: JSON.stringify(data),
      })
      .then((response) => {
        return response.json();
      })
      .then((json) => {
        resolve(json);
      })
      .catch((err) => {
        console.log(err);
      });
  });
  return result;
};
