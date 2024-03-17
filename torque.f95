program main
! The program converts Nm to Kg applied to a shaft of any length
implicit none
use, intrinsic :: iso_fortran_env, only: qp => real128

  real(qp), parameter :: g = 9.80665_qp
  real(qp) :: Length, r, Force, tau

  print '(a)', 'Enter L in cm:'
  read *, Length
  print '(a)', 'Enter desired Nm:'
  read *, Force

  r = Length * 0.01_qp
  tau = Force / r
  print '(a,f6.3,a)', 'Apply ', tau / g, ' kg'
end program main
