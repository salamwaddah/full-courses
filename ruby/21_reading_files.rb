File.open(__dir__ + "/files/employees.txt", "r") do |file|
  for line in file.readlines
    puts line.upcase
  end
end
