declare namespace API {
  type BaseResponseBoolean = {
    code: number
    msg: string
    data: boolean
  }

  type BaseResponseLoginUserResponse = {
    code: number
    msg: string
    data: LoginUserResponse
  }

  type BaseResponseLong = {
    code: number
    msg: string
    data: number
  }

  type BaseResponseString = {
    code: number
    msg: string
    data: string
  }

  type BaseResponseUser = {
    code: number
    msg: string
    data: User
  }

  type BaseResponseUserPage = {
    code: number
    msg: string
    data: UserPage
  }

  type BaseResponseUserResponse = {
    code: number
    msg: string
    data: UserResponse
  }

  type DeleteUserParams = {
    id: string
  }

  type GeneratorRequest = {
    message: string
    /**  Content type: html, multi_file */
    type: string
  }

  type GeneratorResponse = {
    result: string
  }

  type GetUserParams = {
    id: string
  }

  type IdPathRequest = true

  type IdRequest = {
    id: number
  }

  type ListUserByPageParams = {
    current: number
    pageSize: number
    sortField?: string
    sortOrder?: string
    id?: number
    unionId?: string
    mpOpenId?: string
    userName?: string
    userProfile?: string
    userRole?: string
  }

  type LoginUserResponse = {
    id: number
    userAccount: string
    userName: string
    userAvatar: string
    userProfile: string
    userRole: string
    createTime: string
    updateTime: string
    token: string
  }

  type UpdateUserParams = {
    id: string
  }

  type User = {
    id: number
    userAccount: string
    userName: string
    userAvatar: string
    userProfile: string
    userRole: string
    createTime: string
    updateTime: string
  }

  type UserAddRequest = {
    userName?: string
    userAccount: string
    userAvatar?: string
    userProfile?: string
    userRole: string
  }

  type UserLoginRequest = {
    userAccount: string
    userPassword: string
  }

  type UserPage = {
    records: UserResponse[]
    total: number
  }

  type UserQueryRequest = {
    current: number
    pageSize: number
    sortField?: string
    sortOrder?: string
    id?: number
    unionId?: string
    mpOpenId?: string
    userName?: string
    userProfile?: string
    userRole?: string
  }

  type UserRegisterRequest = {
    userAccount: string
    userPassword: string
    checkPassword: string
  }

  type UserResponse = {
    id: number
    userAccount: string
    userName: string
    userAvatar: string
    userProfile: string
    userRole: string
    createTime: string
    updateTime: string
  }

  type UserUpdateRequest = {
    userName?: string
    userAvatar?: string
    userProfile?: string
    userRole?: string
  }
}
