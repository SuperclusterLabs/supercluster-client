import { render, screen } from '@testing-library/react';
import App from './App';

test('Renders Supercluster Files', () => {
  render(<App />);
  const linkElement = screen.getByText(/Supercluster Files/i);
  expect(linkElement).toBeInTheDocument();
});
