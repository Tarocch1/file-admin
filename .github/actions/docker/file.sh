case ${TARGETPLATFORM} in
  linux/amd64) FILENAME="file-admin_linux_64";;
  linux/386) FILENAME="file-admin_linux_32";;
  linux/arm64) FILENAME="file-admin_linux_arm64_v8a";;
  linux/arm/v7) FILENAME="file-admin_linux_arm32_v7a";;
  linux/arm/v6) FILENAME="file-admin_linux_arm32_v6";;
  linux/riscv64) FILENAME="file-admin_linux_riscv64";;
  linux/ppc64le) FILENAME="file-admin_linux_ppc64le";;
  linux/s390x) FILENAME="file-admin_linux_s390x";;
  linux/mips64) FILENAME="file-admin_linux_mips64";;
  linux/mips64le) FILENAME="file-admin_linux_mips64le";;
  *) FILENAME="file-admin_linux_64";;
esac

cp ${FILENAME} file_admin
