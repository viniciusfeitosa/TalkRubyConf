require 'grape'

class API < Grape::API
  params do 
    requires :number, type: Integer
  end
  route_param :number do
    get do
      fib(params[:number])
    end
  end
end

def fib(n)
  if n == 0 || n == 1
    return n
  end
  fib(n-1) + fib(n-2)
end
