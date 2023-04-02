lucky_numbers = [4, 8, 15, 16, 23, 42]

begin
  lucky_numbers["dog"]
  num = 10 / 0
rescue ZeroDivisionError
  puts "Division by zero error"
rescue TypeError => e
  puts e
end
