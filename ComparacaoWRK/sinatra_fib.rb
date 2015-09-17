require 'sinatra'

get '/:number' do
  if params['number'] == "10"
    return "Python + Tornado fib(" + params['number'] + "): " + fib(params['number'].to_i).to_s
  end
end

def fib(n)
  if n == 0 || n == 1
    return n
  end
  fib(n-1) + fib(n-2)
end
