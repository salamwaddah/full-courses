class Student
  attr_accessor :name, :major, :gpa

  def initialize(name, major, gpa)
    @name = name
    @major = major
    @gpa = gpa
  end

  def has_honors
    @gpa > 3.4
  end
end

student1 = Student.new("Salam", "CS", 3.5)
student2 = Student.new("Jim", "Business", 2.4)

puts student1.has_honors
puts student2.has_honors
