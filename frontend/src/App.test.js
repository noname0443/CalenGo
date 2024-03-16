import { render } from '@testing-library/vue'
import App from './App.vue'

test('it should work', () => {
  const { getByText } = render(App, {
    props: {

    }
  })

  getByText('You did it!')
})